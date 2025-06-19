function calculateNumberOfUniqueWaysToClimbAStairs(numberOfSteps) {
  if (numberOfSteps === 0 || numberOfSteps === 1) {
    return 1;
  }

  return (
    calculateNumberOfUniqueWaysToClimbAStairs(numberOfSteps - 1) +
    calculateNumberOfUniqueWaysToClimbAStairs(numberOfSteps - 2)
  );
}

function topDownApproach(numberOfSteps) {
  if (numberOfSteps === 0 || numberOfSteps === 1) {
    return 1;
  }

  let uniqueWays = [1, 1];

  for (let i = 2; i <= numberOfSteps; i++) {
    uniqueWays[i] = uniqueWays[i - 1] + uniqueWays[i - 2];
  }

  return uniqueWays[numberOfSteps];
}

function main() {
  try {
    const numberOfSteps = Number(process.argv[2]);
    const result = topDownApproach(numberOfSteps);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
