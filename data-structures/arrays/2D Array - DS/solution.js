// NOTE to run it
// node solution.js  < input.txt

process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

const getHourglassSum = (array2D, middle) => {
  const { x, y } = middle;

  const value = array2D[y - 1][x - 1] + array2D[y - 1][x] + array2D[y - 1][x + 1]
    + array2D[y][x]
    + array2D[y + 1][x - 1] + array2D[y + 1][x] + array2D[y + 1][x + 1];

  return value;
}

function main() {
  var arr = [];

  for(arr_i = 0; arr_i < 6; arr_i++) {
    arr[arr_i] = readLine().split(' ');
    arr[arr_i] = arr[arr_i].map(Number);
  }

  let maxSum = -1337;

  for (let y = 1 ; y < 5 ; y++) {
    for (let x = 1 ; x < 5 ; x++) {
    const hourGlassValue = getHourglassSum(arr, { x, y });
    maxSum = hourGlassValue > maxSum ? hourGlassValue : maxSum;
    }
  }

  console.log(maxSum);
}
