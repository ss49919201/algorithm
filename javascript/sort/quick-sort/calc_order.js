'use strict';

let order = 0;

const sort = x => {
    if (x.length <= 1) return x;

    const pivotIdx = Math.floor(x.length / 2);
    const left = [];
    const right = [];

    for (let i = 0; i < x.length; i++) {
        if (i === pivotIdx) continue;
        order++

        if (x[pivotIdx] > x[i]) {
            left.push(x[i]);
        } else {
            right.push(x[i]);
        }
    };
    return [...sort(left), x[pivotIdx], ...sort(right)];
};

console.log(sort([]));
console.log(`order: ${order}`);
order = 0;

console.log(sort([1]));
console.log(`order: ${order}`);
order = 0;

console.log(sort([1, 2]));
console.log(`order: ${order}`);
order = 0;

console.log(sort([4, 2]));
console.log(`order: ${order}`);
order = 0;

console.log(sort([1, 4, 0, 3, 3, 2]));
console.log(`order: ${order}`);
order = 0;
