function songPriorityFromPlaylist(rankedList) {
  // While the rankedList is not empty, setup a while loop
  // Initiate a variable (playlistIndex) and firstItemOfPlaylist, and qualified = true that points to the current playlist
  // Go through the other playlist for each of the playlistIndex
  // If firstItemOfPlaylist is in the array and not position 1 then qualified = false
  // At the end of the loop, if qualified is true, add the firstItemOfPlaylist to result and splice the first element from its playlist
  // if qualified === true || playlistIndex is last item, then reset the playlistIndex = 0, else playlistIndex += 1
  // If the playlists are empty then exit the loop

  let playlistIndex = 0;
  const result = [];

  while (true) {
    const playlist = rankedList[playlistIndex];
    const firstSongInPlaylist = playlist[0];

    let isTopPriority = true;

    for (let i = 0; i < rankedList.length; i++) {
      if (i === playlistIndex) {
        continue;
      }

      if (
        rankedList[i].indexOf(firstSongInPlaylist) !== -1 &&
        rankedList[i].indexOf(firstSongInPlaylist) !== 0
      ) {
        isTopPriority = false;
      }
    }

    if (isTopPriority) {
      // Add the song to the list of result
      result.push(firstSongInPlaylist);

      // Remove the song from all playlists
      let i = 0;
      while (rankedList.length !== 0 && i < rankedList.length) {
        let p = rankedList[i];

        if (p.indexOf(firstSongInPlaylist) !== -1) {
          rankedList[i].splice(0, 1);
        }

        if (rankedList[i].length === 0) {
          rankedList.splice(i, 1);
        } else {
          i++;
        }
      }
    }

    if (isTopPriority || playlistIndex === rankedList.length - 1) {
      playlistIndex = 0;
    } else {
      playlistIndex++;
    }

    if (rankedList.length === 0) {
      break;
    }
  }

  return result;
}

function main() {
  try {
    const rankedList = JSON.parse(process.argv[2]);

    const result = songPriorityFromPlaylist(rankedList);
    console.log("The result is ", result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
