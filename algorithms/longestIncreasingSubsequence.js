function longestIncreasingSubsequence(array, longestSubsequence) {
  if (array.length === 0) {
    return longestSubsequence;
  }

  let result = [...longestSubsequence];

  for (let i = 0; i < array.length; i++) {
    const currentNumber = array[i];
    const lastNumberInLongestSubsequence =
      longestSubsequence[longestSubsequence.length - 1];

    const isLongestSubSequenceStillValid =
      lastNumberInLongestSubsequence === undefined ||
      currentNumber > lastNumberInLongestSubsequence;

    if (isLongestSubSequenceStillValid) {
      const newLongestSubsequence = longestIncreasingSubsequence(
        array.slice(i + 1),
        [...longestSubsequence, currentNumber]
      );

      if (newLongestSubsequence.length > result.length) {
        result = [...newLongestSubsequence];
      }
    }
  }

  return result;
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]);
    const result = longestIncreasingSubsequence(array, []);

    console.log(result);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
