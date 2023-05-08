main();

function main() {
  const arr = [2, 3, 10, 0, 1, 6];
  const want = 100;
  for (const v of arr) {
    console.debug(v);
    if (v === want) {
      console.info("Hit!");
      return
    }
  }
  console.info("Not found...")
}
