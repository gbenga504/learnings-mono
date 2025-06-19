function largestSumOfNonAdjacentNumbers(array) {
  let maximumSumWithCurrentNumberInclusive = 0;
  let maximumSumWithCurrentNumberExclusive = 0;

  array.forEach((number, index) => {
    if (index === 0) {
      maximumSumWithCurrentNumberExclusive = 0;
      maximumSumWithCurrentNumberInclusive = number;
    } else {
      let temporaryMaximumSumWithCurrentNumberExclusive =
        maximumSumWithCurrentNumberExclusive;

      maximumSumWithCurrentNumberExclusive =
        maximumSumWithCurrentNumberInclusive;

      maximumSumWithCurrentNumberInclusive = Math.max(
        number + temporaryMaximumSumWithCurrentNumberExclusive,
        maximumSumWithCurrentNumberInclusive
      );
    }
  });

  return Math.max(
    maximumSumWithCurrentNumberExclusive,
    maximumSumWithCurrentNumberInclusive
  );
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    const result = largestSumOfNonAdjacentNumbers(array);

    console.log(result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
