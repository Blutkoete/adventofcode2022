const fs = require('fs');
const allContents = fs.readFileSync('input.txt', 'utf-8');
const lines = allContents.split(/\r?\n/)
const pairs = []
const packets = []
for (var i = 0; i < lines.length; i = i + 3) {
    const left = lines[i]
    const right = lines[i + 1]

    pairs.push([JSON.parse(left), JSON.parse(right)])
    packets.push(JSON.parse(left))
    packets.push(JSON.parse(right))
}

function compareArrays(left, right) {
    for (var i = 0; i < left.length; i++) {
        const leftElement = left[i]
        if (right.length == i) {
            console.log("Right side ran out of elements")
            return -1
        }
        const rightElement = right[i]
        const leftType = typeof leftElement
        const rightType = typeof rightElement
        if (leftType == "number" && rightType == "number") {
            if (leftElement > rightElement) {
                return -1
            } else if (leftElement < rightElement) {
                return 1
            }
        } else if (leftType == "object" && rightType == "number") {
            const res = compareArrays(leftElement, [rightElement])
            if (res == -1) {
                return -1
            } else if (res == 1) {
                return 1
            }
        } else if (leftType == "number" && rightType == "object") {
            const res = compareArrays([leftElement], rightElement)
            if (res == -1) {
                return -1
            } else if (res == 1) {
                return 1
            }
        } else if (leftType == "object" && rightType == "object") {
            const res = compareArrays(leftElement, rightElement)
            if (res == -1) {
                return -1
            } else if (res == 1) {
                return 1
            }
        }
    }
    if (left.length == i && right.length > i) {
        console.log("Left side ran out of elements")
        return 1
    }
    return 0
}


var i = 1
const rightPairs = []
pairs.forEach(function (pair) {
    left = pair[0]
    right = pair[1]
    console.log("Checking pair ", pair)
    const result = compareArrays(left, right)
    console.log("result ", result)
    if (result == 1) {
        rightPairs.push(i)
    }
    i++
})

console.log("part 1", rightPairs.reduce((a, b) => a + b, 0))

packets.push([[2]])
packets.push([[6]])
packets.sort(function (a, b) { return compareArrays(b, a) })

decoderKey = 1
packets.forEach((val, index) => {
    if (JSON.stringify(val) == JSON.stringify([[2]])) {
        console.log("Found 2", index)
        decoderKey = decoderKey * (index + 1)
    } else if (JSON.stringify(val) == JSON.stringify([[6]])) {
        console.log("Found 6", index)
        decoderKey = decoderKey * (index + 1)
    }
})
console.log("part 2", decoderKey)