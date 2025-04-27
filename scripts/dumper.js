const { MongoClient, ServerApiVersion } = require('mongodb');

// Get the MONGO_URL environment variable
const mongoUrl = process.env.MONGO_URI;
// The name of the database and collection to use
const dbName = process.env.MONGO_DATABASE || "vtubers";
const collectionName = process.env.MONGO_COLLECTION || "vtubers";

// Create a MongoClient with a MongoClientOptions object to set the Stable API version
const client = new MongoClient(mongoUrl,  {
    serverApi: {
        version: ServerApiVersion.v1,
        strict: true,
        deprecationErrors: true,
        setTimeout: 10000,
    }
});

console.log(`Connecting to ${mongoUrl}`);

async function run() {
  try {
    const database = client.db(dbName);
    const collection = database.collection(collectionName);

    // Fetch all documents from the collection
    const cursor = collection.find({});
    const documents = await cursor.toArray();
    console.log(`Found ${documents.length} documents in the collection.`);
    // Output the documents to the console
    documents.forEach(doc => {
      console.log(doc);
    });
    console.log("Dump completed successfully.");
    
  } finally {
    // Close the MongoDB client connection
    await client.close();
  }
}

// Run the function and handle any errors
run().catch(console.dir);
