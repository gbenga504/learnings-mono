function largestSumofSubArrayBruteForce(array) {
  let largestSum = 0;

  for (let i = 0; i <= array.length - 1; i++) {
    let currentSum = 0;

    for (let j = i; j <= array.length - 1; j++) {
      currentSum = array[j] + currentSum;
      if (currentSum > largestSum) {
        largestSum = currentSum;
      }
    }
  }

  return largestSum;
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    const result = largestSumofSubArrayBruteForce(array);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
