function findFirstMissingPositiveInteger(array) {
  return array.reduce((acc, number) => {
    let lowerBoundary = number - 1;
    let upperBoundary = number + 1;

    if (
      array.indexOf(lowerBoundary) === -1 &&
      lowerBoundary <= acc &&
      lowerBoundary > 0
    ) {
      acc = lowerBoundary;
    }

    if (
      array.indexOf(upperBoundary) === -1 &&
      upperBoundary <= acc &&
      upperBoundary > 0
    ) {
      acc = upperBoundary;
    }

    return acc;
  }, Number.MAX_SAFE_INTEGER);
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    const result = findFirstMissingPositiveInteger(array);

    console.log(result);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
