package chicken

import (
	"context"
	"libro-electronico/model"
	"net"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// MongoConnect establishes a connection to the MongoDB database.
func MongoConnect(mconn model.DBIngfo) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		mconn.DBString = SRVLookup(mconn.DBString)
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
		if err != nil {
			return nil, err
		}
	}
	db = client.Database(mconn.DBName)
	return db, nil
}

// SRVLookup performs an SRV lookup and returns the modified MongoDB URI.
func SRVLookup(srvuri string) (mongouri string) {
	// Split the SRV URI to extract user/pass and domain/dbname
	atsplits := strings.Split(srvuri, "@")
	userpass := strings.Split(atsplits[0], "//")[1]
	mongouri = "mongodb://" + userpass + "@"
	slashsplits := strings.Split(atsplits[1], "/")
	domain := slashsplits[0]
	dbname := slashsplits[1]

	// Set up DNS resolver
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 10,
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	// Perform SRV lookup
	_, srvs, err := r.LookupSRV(context.Background(), "mongodb", "tcp", domain)
	if err != nil {
		panic("SRV lookup failed: " + err.Error())
	}

	// Build the SRV list
	var srvlist string
	for _, srv := range srvs {
		srvlist += strings.TrimSuffix(srv.Target, ".") + ":" + strconv.FormatUint(uint64(srv.Port), 10) + ","
	}

	// Lookup TXT records
	txtrecords, err := r.LookupTXT(context.Background(), domain)
	if err != nil {
		panic("TXT lookup failed: " + err.Error())
	}
	var txtlist string
	for _, txt := range txtrecords {
		txtlist += txt + "&"
	}

	// Construct the MongoDB URI
	mongouri = mongouri + strings.TrimSuffix(srvlist, ",") + "/" + dbname + "?ssl=true&" + strings.TrimSuffix(txtlist, "&")
	return mongouri
}


func DeleteOneDoc(db *mongo.Database, collection string, filter bson.M) (updateresult *mongo.DeleteResult, err error) {
	updateresult, err = db.Collection(collection).DeleteOne(context.Background(), filter)
	return
}

func DeleteManyDocs(db *mongo.Database, collection string, filter bson.M) (deleteresult *mongo.DeleteResult, err error) {
	deleteresult, err = db.Collection(collection).DeleteMany(context.Background(), filter)
	return
}

func GetAllDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		return
	}
	return
}

func GetAllDistinctDoc(db *mongo.Database, filter bson.M, fieldname, collection string) (doc []any, err error) {
	ctx := context.TODO()
	doc, err = db.Collection(collection).Distinct(ctx, fieldname, filter)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	err = db.Collection(collection).FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		return
	}
	return
}

func GetOneLatestDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	err = db.Collection(collection).FindOne(context.TODO(), filter, opts).Decode(&doc)
	if err != nil {
		return
	}
	return
}

// FindDocs mencari dokumen dalam koleksi berdasarkan filter yang diberikan
func FindDocs(database *mongo.Database, collection string, filter bson.M) (*mongo.Cursor, error) {
	// Membuat context dengan timeout 10 detik
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mengakses koleksi yang diinginkan
	coll := database.Collection(collection)

	// Membuat opsi pencarian (misalnya, untuk mengatur batasan hasil, mengurutkan, dll)
	opts := options.Find()

	// Melakukan pencarian dokumen dengan filter yang diberikan
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func CountDocs(db *mongo.Database, collection string, filter bson.M) (count int64, err error) {
	count, err = db.Collection(collection).CountDocuments(context.Background(), filter)
	if err != nil {
		return
	}
	return
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}, err error) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return
	}
	return insertResult.InsertedID, nil
}

func UpdateOneDoc(db *mongo.Database, collection string, filter bson.M, updatefield bson.M) (updateresult *mongo.UpdateResult, err error) {
	updateresult, err = db.Collection(collection).UpdateOne(context.TODO(), filter, updatefield)
	if err != nil {
		return
	}
	return
}

func ReplaceOneDoc(db *mongo.Database, collection string, filter bson.M, doc interface{}) (updateresult *mongo.UpdateResult, err error) {
	updateresult, err = db.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		return
	}
	return
}

func GetOneLowestDoc[T any](db *mongo.Database, collection string, filter bson.M, sortField string) (doc T, err error) {
	opts := options.FindOne().SetSort(bson.M{sortField: 1}) // Sort by the provided field in ascending order
	err = db.Collection(collection).FindOne(context.TODO(), filter, opts).Decode(&doc)
	if err != nil {
		return
	}
	return
}