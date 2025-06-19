function areaOfRectangleFromHistogram(histogram) {
  const sortedHistogram = histogram.sort((a, b) => a - b);
  let result = 0;

  for (let i = 0; i < histogram.length; i++) {
    // height * breadth
    const area = sortedHistogram[i] * (histogram.length - i);

    if (area > result) {
      result = area;
    }
  }

  return result;
}

function main() {
  try {
    const histogram = JSON.parse(process.argv[2]);

    const result = areaOfRectangleFromHistogram(histogram);
    console.log("The result is ", result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
