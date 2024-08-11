// Part1
const fs = require("fs");

let commandList = [];
fs.readFileSync("data.input", "utf-8").split(/\r?\n/).forEach(line => {

    let splittedLine = line.trim().split(/ /);
    commandList.push({
        dir: splittedLine[0].toLowerCase(),
        val: Number(splittedLine[1])
    });
});

let pos = {horizontal: 0, depth: 0};
commandList.forEach(cmd => {
    if (cmd.dir == "forward") {
        pos.horizontal += cmd.val;
    } else if (cmd.dir == "up") {
        pos.depth -= cmd.val;
    } else if (cmd.dir == "down") {
        pos.depth += cmd.val;
    } else {
        console.log(`unknown direction: ${cmd.dir}`);
    }
});

console.log(``);
console.log(`final submarine position: ${pos.horizontal}, ${pos.depth}`);
console.log(`puzzle 1 solution: ${(pos.horizontal * pos.depth)}`);

// Part2

let sub = {aim: 0, horizontal: 0, depth: 0};
commandList.forEach(cmd => {
    if (cmd.dir == "forward") {
        sub.horizontal += cmd.val;
        sub.depth += cmd.val * sub.aim;
    } else if (cmd.dir == "up") {
        sub.aim -= cmd.val;
    } else if (cmd.dir == "down") {
        sub.aim += cmd.val;
    } else {
        console.log(`unknown direction: ${cmd.dir}`);
    }
});

console.log(``);
console.log(`final sub position: ${sub.horizontal}, ${sub.depth}`);
console.log(`puzzle 2 solution: ${(sub.horizontal * sub.depth)}`);