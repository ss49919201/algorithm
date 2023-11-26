function search(arr, tgt) {
    if (arr.length === 0) {
        return false;
    }
 
    if (arr.length === 1) {
        return arr[0] === tgt;
    }

    const mid = Math.floor(arr.length / 2)
    if (arr[mid] === tgt) {
        return true;
    }

    if (arr[mid] > tgt) {
        return search(arr.slice(0, mid-1), tgt);
    } else {
        return search(arr.slice(mid), tgt);
    }
};

console.log(search([1, 3, 100, 1_000], 1_000));
console.log(search([1, 3], 2));
console.log(search([1], 2));
console.log(search([], 2));
