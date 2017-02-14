'use strict';
//---------------------Classes and Function - to be moved later --------------//

var removeDataColors = function (data) {
    data.knownTech = [];
    for (var i in data.languages) {
        data.knownTech.push(data.languages[i].name);
    }
    //data.knownTech.sort();
};

// There should be a better way to do this, but I am blanking now -- maybe filter
// Corrects data formatting
var rowClass = function (data, key) {
    data.name = data.firstName + " " + data.lastName;
    data.namelink = '<a ui-sref="profile/' + key + '">' + data.name + "</a>";
    data.gradeLabel = data.grade;
    data.key = key;
    removeDataColors(data);

    //console.log(data);
    return data;
};

var removedDuplicates = function (arr) {
    return arr.filter(function (elem, index, self) {
        return index == self.indexOf(elem);
    });
};

var cleanStudents = function (student) {
    //student.interestedIn = removedDuplicates(student.interestedIn);
    //student.languages = removedDuplicates(student.interestedIn);
    return student;
};
