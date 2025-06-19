let count = 0;

function decodeMessage(encodedMessage) {
  let mapping = 1;

  while (mapping <= 26) {
    let digit = Number(encodedMessage.substring(0, 1));
    let tens = Number(encodedMessage.substring(0, 2));

    if (digit === mapping && encodedMessage.length === 1) {
      count++;
      break;
    } else if (tens === mapping && encodedMessage.length === 2) {
      count++;
      break;
    } else if (digit === mapping && encodedMessage.length !== 1) {
      decodeMessage(encodedMessage.substring(1));
    } else if (tens === mapping && encodedMessage.length !== 2) {
      decodeMessage(encodedMessage.substring(2));
    }

    mapping++;
  }
}

function main() {
  try {
    const encodedMessage = process.argv[2];
    decodeMessage(encodedMessage);
    console.log("The result is ", count);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
