function hIndex(citationsInPapers) {
  let result = 0;

  for (let i = 0; i < citationsInPapers.length; i++) {
    let currentCitation = citationsInPapers[i];

    let papersWithCitations = citationsInPapers.filter(
      (c) => c >= currentCitation
    );

    if (
      papersWithCitations.length === currentCitation &&
      currentCitation > result
    ) {
      result = currentCitation;
    }
  }

  return result;
}

function main() {
  try {
    const citationsInPapers = JSON.parse(process.argv[2]);
    const result = hIndex(citationsInPapers);

    console.log(result);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
