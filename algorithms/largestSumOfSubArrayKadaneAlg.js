function largestSumofSubArrayKadaneAlg(array) {
  let globalLargestSum = array[0];
  let currentLargestSum = globalLargestSum;

  for (let i = 1; i <= array.length - 1; i++) {
    currentLargestSum = Math.max(array[i], array[i] + currentLargestSum);

    if (currentLargestSum > globalLargestSum) {
      globalLargestSum = currentLargestSum;
    }
  }

  return globalLargestSum;
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    const result = largestSumofSubArrayKadaneAlg(array);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
