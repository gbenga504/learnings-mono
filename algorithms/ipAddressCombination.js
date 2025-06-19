function ipAddressCombination({ remainingDigits, computedDigits, result }) {
  // We will be using a Trie like data structure to solve this
  // Loop through the remaining digits from [index 0  ... increment]
  // compute the new computedDigits, attachments etc
  // If the conditions are met, then call the function [ipAddressCombination] again with the new values
  // If everything works well then push into the result and return
  for (let i = 0; i <= remainingDigits.length - 1; i++) {
    let digitsToAppend = remainingDigits.substring(0, i + 1);

    if (digitsToAppend.length > 1 && digitsToAppend[0] === "0") {
      break;
    }

    digitsToAppend = Number(digitsToAppend);

    if (digitsToAppend > 255) {
      break;
    }

    const seperator = computedDigits.length === 0 ? "" : ".";

    const newComputedDigits = `${computedDigits}${seperator}${digitsToAppend}`;
    const newRemainingDigits = remainingDigits.substring(i + 1);

    if (newComputedDigits.split(".").length > 4) {
      break;
    }

    if (
      newComputedDigits.split(".").length === 4 &&
      i === remainingDigits.length - 1
    ) {
      result.push(newComputedDigits);

      break;
    }

    if (i !== remainingDigits.length - 1) {
      ipAddressCombination({
        remainingDigits: newRemainingDigits,
        computedDigits: newComputedDigits,
        result,
      });
    }
  }

  return result;
}

function main() {
  // node ipAddressCombination "2542540123"
  // The result is ===> [ '254.25.40.123', '254.254.0.123' ]
  try {
    const digits = process.argv[2];
    const result = ipAddressCombination({
      remainingDigits: digits,
      computedDigits: "",
      result: [],
    });

    console.log("The result is ===>", result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
