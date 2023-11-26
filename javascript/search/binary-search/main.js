function search(arr, tgt, start, end) {
    if (start > end) return false;
 
    const mid = Math.floor((start+end) / 2)
    if (arr[mid] === tgt) return true;

    if (arr[mid] > tgt) {
        return search(arr, tgt, start, mid-1);
    } else {
        return search(arr, tgt, mid+1, end);
    }
};

console.log(search([1, 3, 100, 1_000], 1_000, 0, 3));
console.log(search([1, 3, 100, 1_000], 1, 0, 3));
console.log(search([1, 3], 2, 0, 1));
console.log(search([1], 2, 0, 0));
console.log(search([], 2, 0, 0));
