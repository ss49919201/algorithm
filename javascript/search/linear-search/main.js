main();

function main() {
  const arr = [2, 3, 10, 0, 1, 6],
        want = 100;
  console.info(search(arr, want)); // false

  const want2 = 3;
  console.info(search(arr, want2)); // true
}

function search(arr, want) {
  for (const v of arr) {
    if (v === want) {
      return true;
    }
  }
  return false;
}