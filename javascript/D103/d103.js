process.stdin.resume();
process.stdin.setEncoding('utf8');
var lines = [];
var reader = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});
reader.on('line', (line) => {
  lines.push(line);
});
reader.on('close', () => {
        
  const split = lines[0].split("");
  let output = "";
  for (i=split.length - 1; i >= 0; i--) {
      output += split[i]
  }

  console.log(output);
});