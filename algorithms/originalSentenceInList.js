function originalSentenceInList({
  dictionary,
  sentence,
  matches,
  indexOfMatch,
}) {
  // This is the base case, when the sentence is of length 0, then there is nothing to do
  if (sentence.length === 0) return;

  let endIndex = 0;

  //We take substrings and check if its in the dictionary. If it is then we recurse over the new string
  while (endIndex < sentence.length) {
    const currentStringToCheck = sentence.substring(0, endIndex + 1);

    if (dictionary.indexOf(currentStringToCheck) !== -1) {
      if (indexOfMatch === undefined) {
        matches.push([currentStringToCheck]);
      } else {
        matches[indexOfMatch].push(currentStringToCheck);
      }

      originalSentenceInList({
        dictionary,
        sentence: sentence.replace(currentStringToCheck, ""),
        matches,
        indexOfMatch: indexOfMatch ?? matches.length - 1,
      });
    }

    endIndex++;
  }
}

function main() {
  try {
    const dictionary = JSON.parse(process.argv[2]);
    const sentence = process.argv[3];
    const matches = [];

    originalSentenceInList({
      dictionary,
      sentence,
      matches,
    });

    console.log(matches.length > 0 ? matches : null);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
