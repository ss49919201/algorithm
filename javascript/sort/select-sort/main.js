'use strict';

const sort = arr => {
    let result = [];
    while (arr.length > 0) {
        let min = arr[0];
        let minIdx = 0;
        for (let i = 1; i < arr.length; i++) {
            if (min > arr[i]) {
                min = arr[i];
                minIdx = i;
            }
        };
        result.push(arr[minIdx]);
        arr = [...arr.slice(0, minIdx), ...arr.slice(minIdx+1)];
    }
    return result;
};

console.log(sort([]));
console.log(sort([1]));
console.log(sort([3, 1, 0, 2, -1]));
