function minimumNumberOfClassroomsRequired(timeIntervals) {
  let ongoingLectures = [];

  // If the time interval only contains a single slot then we only need one classroom
  if (timeIntervals.length === 1) {
    return 1;
  }

  // We need to sort the time intervals by ascending starting time slots
  timeIntervals.sort((a, b) => a[0] - b[0]);

  //We check for the next available slots
  timeIntervals.forEach((timeInterval) => {
    const indexOfNextAvailableSlot = ongoingLectures.findIndex(
      (ongoingLecture) => timeInterval[0] >= ongoingLecture[1]
    );

    if (indexOfNextAvailableSlot >= 0) {
      ongoingLectures.splice(indexOfNextAvailableSlot, 1, timeInterval);
    } else {
      ongoingLectures.push(timeInterval);
    }
  });

  return ongoingLectures.length;
}

function main() {
  try {
    const timeIntervals = JSON.parse(process.argv[2]);
    const result = minimumNumberOfClassroomsRequired(timeIntervals);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
