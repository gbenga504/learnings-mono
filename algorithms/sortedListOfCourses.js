function sortedListOfCourses(courseIdHashMap) {
  const result = [];
  const totalNumberOfCourseIds = Object.keys(courseIdHashMap).length;

  for (let i = 0; i < totalNumberOfCourseIds; i++) {
    // [i] represents the dependent that a course has.
    // E.g if i === 0, we need to look for a course that does not depend on a completion of any course
    // if i === 1, we need to look for a course that depends on exactly 1 course
    const courseIds = Object.keys(courseIdHashMap).filter(
      (cId) => courseIdHashMap[cId].length === i
    );

    // At every stage we should only have 1 course bcos every class/course has a unique number of dependents
    if (courseIds.length !== 1) {
      break;
    }

    const courseId = courseIds[0];
    const isNumberOfDependentCourseIdsEqualToRegistered =
      courseIdHashMap[courseId].length === result.length;

    const isDependentCourseIdsRegistered = courseIdHashMap[courseId].reduce(
      (acc, cId) => {
        return acc && result.indexOf(cId) !== -1;
      },
      isNumberOfDependentCourseIdsEqualToRegistered
    );

    // If the dependent courses are not registered yet then it means there is an issue with the courses
    if (!isDependentCourseIdsRegistered) {
      break;
    } else {
      result.push(courseId);
    }
  }

  if (result.length !== totalNumberOfCourseIds) {
    return null;
  }

  return result;
}

// node sortedListOfCourses '{"CSC300":["CSC100", "CSC200"],"CSC200":["CSC100"],"CSC100":[]}'
// The result is  [ 'CSC100', 'CSC200', 'CSC300' ]
function main() {
  try {
    const courseIdHashMap = JSON.parse(process.argv[2]);

    const result = sortedListOfCourses(courseIdHashMap);
    console.log("The result is ", result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
