const fs = require('fs')
const path = require('path');
const { MongoClient, ServerApiVersion } = require('mongodb');
const { parse } = require('csv-parse/sync');
const { generateKeyPair } = require('crypto');

// Get the MONGO_URL environment variable
const mongoUrl = process.env.MONGO_URI;

const dbName = "vtubers";
const collectionName = "vtubers";

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

// Create a MongoClient with a MongoClientOptions object to set the Stable API version
const client = new MongoClient(mongoUrl,  {
    serverApi: {
        version: ServerApiVersion.v1,
        strict: true,
        deprecationErrors: true,
        setTimeout: 10000,
    }
});

async function run() {
  try {
    const database = client.db(dbName);
    const collection = database.collection(collectionName);

    vtuberFile = '/data/vtuber.csv';
    vtuberData = loadCSV(vtuberFile)
    vtuber = fileToMap(vtuberFile);

    generationFile = '/data/generation.csv';
    generation = fileToMap(generationFile)

    companyFile = '/data/company.csv';
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
