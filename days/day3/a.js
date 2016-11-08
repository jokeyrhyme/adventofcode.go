'use strict'

const fs = require('fs')
const path = require('path')

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf8')

const houses = {}

let ns = 0
let ew = 0

let coords = `${ns},${ew}`
houses[coords] = true

input.split('').forEach((move) => {
  if (move == "^") {
    ns += 1
  }
  if (move == "v") {
    ns -= 1
  }
  if (move == ">") {
    ew += 1
  }
  if (move == "<") {
    ew -= 1
  }
  coords = `${ns},${ew}`
  houses[coords] = true
})

const result = Object.keys(houses).length

console.log(`day 3: A: ${result}`)
