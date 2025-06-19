function calculateLongestAbsolutePathToAFile(
  directoryStructure,
  directoryMap,
  longestAbsolutePathToAFile
) {
  //Handle the case where the user passes a root directory structure that only contains dir
  if (directoryStructure === "dir") {
    return longestAbsolutePathToAFile;
  }

  let currentFileOrDirectory = getCurrentDirectoryOrFile(directoryStructure);

  // Check if this is a file and hence also check if the filePath is the longest and assign it to
  // longestAbsolutePathToAFile if true
  if (currentFileOrDirectory.match(/.+\d*\.ext/)) {
    const filePath = getFilePath(directoryMap, currentFileOrDirectory);

    if (filePath.length >= longestAbsolutePathToAFile.length) {
      longestAbsolutePathToAFile = filePath;
    }
  }

  directoryStructure = directoryStructure.replace(currentFileOrDirectory, "");
  const nextDepthMatch = directoryStructure.match(/\n(\t)+/);

  if (nextDepthMatch) {
    // If we have a depth match then we wanna remove the depth from the directory structure
    // We also want to add the current directory to our directory map if it does not already
    // exist in the map

    //When moving up the directories i.e backtracking, we wanna remove the old directory as it is stale
    const nextDepth = nextDepthMatch[0].split("\t").length - 1;
    directoryStructure = directoryStructure.replace(nextDepthMatch[0], "");

    if (directoryMap.has(nextDepth + 1)) {
      directoryMap.delete(nextDepth + 1);
    }

    if (!directoryMap.has(nextDepth)) {
      directoryMap.set(nextDepth, currentFileOrDirectory);
    }

    return calculateLongestAbsolutePathToAFile(
      directoryStructure,
      directoryMap,
      longestAbsolutePathToAFile
    );
  } else {
    // Handle the base case where we are done going through the directory structure. Here, we only wanna return the result so
    // it bubbles up through the recursive stack back to the caller
    return longestAbsolutePathToAFile;
  }
}

function getCurrentDirectoryOrFile(directoryStructure) {
  const fileOrDirectoryMatch = directoryStructure.match(/.+(?=\n(\t)+)/);

  // When the file or directory match is null and the directory structure is not empty, then
  // the current directory or file is the directory structure
  if (fileOrDirectoryMatch === null && directoryStructure.length > 0) {
    return directoryStructure;
  }

  return fileOrDirectoryMatch[0];
}

function getFilePath(directoryMap, currentFile) {
  let path = "";

  directoryMap.forEach((directory) => {
    path += `${directory}\\`;
  });

  return `${path}${currentFile}`;
}

function main() {
  // "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
  try {
    const directoryStructure = process.argv[2];
    const directoryMap = new Map();

    const result = calculateLongestAbsolutePathToAFile(
      directoryStructure,
      directoryMap,
      ""
    );

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
