const fs = require("fs");
const path = require("path");

const directoryPath = path.join(__dirname, "../../../internal/");
const modules = [];
try {
  const files = fs.readdirSync(directoryPath);
  files.forEach((file) => {
    const filePath = path.join(directoryPath, file);
    const stat = fs.statSync(filePath);
    if (stat.isDirectory()) {
      modules.push({
        name: file,
        value: file,
      });
    }
  });
} catch (err) {
  console.log("Unable to scan directory: " + err);
}

module.exports = [
  {
    type: "select",
    name: "module",
    message: "Em qual módulo devo criar o usecase",
    choices: modules,
  },
  {
    type: 'input',
    name: 'name',
    message: "What is the use case name"
  }
]
