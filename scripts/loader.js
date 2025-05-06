const fs = require('fs');
const path = require('path');
const { MongoClient, ServerApiVersion } = require('mongodb');
const { parse } = require('csv-parse/sync');

// Get the MONGO_URL environment variable
const mongoUrl = process.env.MONGO_URI;
// The paths to the directory where csv files are stored.
const dataDir = process.env.DATA_DIR;
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

function loadCSV(filePath) {
  let fileData = fs.readFileSync(filePath, 'utf8');
  let records = parse(fileData, {
    cast: function(value, context){
      if (!isNaN(value)) {
        return parseInt(value);
      }
      return value;
    },
    columns: true,
    skip_empty_lines: true,
    trim: true
  });
  return records;
}

function fileToMap(file) {
  dataFromFile = loadCSV(file);
  return new Map(dataFromFile.map(record => { return [record["Id"], record["Name"]]}))
}

async function run() {
  try {
    const database = client.db(dbName);
    const collection = database.collection(collectionName);

    vtuberFile = `${dataDir}/vtuber.csv`;
    vtuberData = loadCSV(vtuberFile)
    vtuber = fileToMap(vtuberFile);

    generationFile = `${dataDir}/generation.csv`;
    generation = fileToMap(generationFile)

    companyFile = `${dataDir}/company.csv`;
    company = fileToMap(companyFile);

    const newRecords = vtuberData.map(record => {
      return {
        "Id": record["Id"],
        "Name": vtuber.get(parseInt(record["Id"])),
        "Generation": generation.get(parseInt(record["GenerationId"])),
        "Company": company.get(parseInt(record["CompanyId"])),
      }
    })

    // Clear out the collection.
    const clearResult = await collection.deleteMany({});

    const result = await collection.insertMany(newRecords);
    
    // Print the ID of the inserted document
    console.log(`A document was inserted with the _id: ${result}`);
  } finally {
    // Close the MongoDB client connection
    await client.close();
  }
}

// Run the function and handle any errors
run().catch(console.dir);
