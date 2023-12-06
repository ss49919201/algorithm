const _ = require('lodash');

const sort = (n, arr) => {
    const m = {};
    for (let i = 0; i <= n; i++) {
        m[i] = 0
    };
    for (const v of arr) {
        m[v]++
    };

    const r = [];
    for (let i = 0; i <= n; i++) {
        for (let j = 0; j < m[i]; j++) {
            r.push(i)
        }
    };
    return r;
};

console.log(_.isEqual(sort(1, [1]), [1]));
console.log(_.isEqual(sort(4, [3, 1, 0, 2, 1]), [0, 1, 1, 2, 3]));
console.log(_.isEqual(sort(3, [2, 1, 0, 2, 1]), [0, 1, 1, 2, 2]));
