const fs = require('fs');
const path = require('path');

const fileCompany = "company.csv";
const fileGeneration = "generation.csv";
const fileVTuber = "vtuber.csv"

const filePath = path.normalize(path.join(__dirname, "../../data", fileCompany));
console.log(`Reading file at ${filePath}`);

const contents = fs.readFileSync(filePath, 'utf8');

console.log(contents);