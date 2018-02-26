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

const pad = x => parseInt(x, 10) < 10 ? `0${parseInt(x, 10)}` : x;

/////////////// ignore above this line ////////////////////

function timeConversion(s) {
    // Complete this function
    let splitted = s.split(':')
    splitted[2] = splitted[2].slice(0, 2);

    let hh = parseInt(splitted[0]);
    const mm = parseInt(splitted[1]);
    const ss = parseInt(splitted[2]);

    const PM = s[s.length - 2] === 'P';

    if (!PM && hh === 12) hh = 0;
    if (PM && hh < 12) hh += 12;

    return `${pad(hh)}:${pad(mm)}:${pad(ss)}`;
}

function main() {
    var s = readLine();
    var result = timeConversion(s);
    process.stdout.write("" + result + "\n");

}
