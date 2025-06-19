function lastSurvivorPrisoner(numberOfPrisoners, numberOfPersonToRemove) {
  const result = [];
  let prisoners = [];

  for (let i = 1; i <= numberOfPrisoners; i++) {
    prisoners.push(i);
  }

  let index = numberOfPersonToRemove - 1;
  let tempPrisonerFromEachPass = [...prisoners];

  while (result.length !== numberOfPrisoners) {
    if (index > prisoners.length - 1) {
      index = index - prisoners.length;

      index = tempPrisonerFromEachPass.length === 1 ? 0 : index;
      prisoners = [...tempPrisonerFromEachPass];
      tempPrisonerFromEachPass = [...prisoners];
    }

    result.push(prisoners[index]);

    // Remove from the tempPrisoner array
    const prionerRemoved = prisoners[index];
    const indexOfPrisonerInTemp =
      tempPrisonerFromEachPass.indexOf(prionerRemoved);
    tempPrisonerFromEachPass.splice(indexOfPrisonerInTemp, 1);

    index = index + numberOfPersonToRemove;
  }

  console.log("gad re", result);
  return result[result.length - 1];
}

function main() {
  try {
    const numberOfPrisoners = Number(process.argv[2]);
    const numberOfPersonToRemove = Number(process.argv[3]);

    const result = lastSurvivorPrisoner(
      numberOfPrisoners,
      numberOfPersonToRemove
    );
    console.log("The result is ", result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
