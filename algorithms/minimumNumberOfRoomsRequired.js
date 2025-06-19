function minimumNumberOfRoomsRequired(lectures) {
  //  [[0, 60], [0, 50], [60, 100], [0, 30]] = 3

  // Init an ongoing lectures array
  // Initiate a loop through the array of lectures
  // If the ongoing lectures is empty then add the lecture into the ongoing lecture
  // Else check if the current lecture overlaps with any ongoing lecture [loop]
  // If overlap === true, then save the index to replace
  // After the inner loop, if there is no overlap, then we need a room; hence add the lecture to the list of ongoing
  // Else replace the index using the splice method in the ongoing lectures array

  let ongoingLectures = [];

  for (let i = 0; i < lectures.length; i++) {
    const lectureToAssign = lectures[i];

    if (ongoingLectures.length === 0) {
      ongoingLectures.push(lectureToAssign);

      continue;
    }

    let lectureToAssignMetadata = { foundASlot: false, slotIndex: null };

    for (let j = 0; j < ongoingLectures.length; j++) {
      const [_, endTimeForOngoingLecture] = ongoingLectures[j];
      const [startTimeForLectureToAssign, endTimeForLectureToAssign] =
        lectureToAssign;

      if (startTimeForLectureToAssign >= endTimeForOngoingLecture) {
        lectureToAssignMetadata = { foundASlot: true, slotIndex: j };

        break;
      }
    }

    if (lectureToAssignMetadata.foundASlot) {
      ongoingLectures.splice(
        lectureToAssignMetadata.slotIndex,
        1,
        lectureToAssign
      );
    } else {
      ongoingLectures.push(lectureToAssign);
    }
  }

  return ongoingLectures.length;
}

function main() {
  try {
    const lectures = JSON.parse(process.argv[2]);
    const result = minimumNumberOfRoomsRequired(lectures);

    console.log("The result is ===>", result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
