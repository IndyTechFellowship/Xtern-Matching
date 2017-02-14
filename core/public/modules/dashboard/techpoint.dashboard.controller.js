'use strict';
angular.module('Xtern')
<<<<<<< HEAD
    .controller('TechPointDashboardCtrl', ['$scope', function ($scope) {
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart: false,
                title: "Universities",
                icon: 'university',
                dataLabel: 'university',
                nestedData: false
            },
            status: {
                isChart: true,
                title: "Stage",
                labels: ["Stage 1 Approved", "Remaining", "Denied"],
                dataLabel: 'status',
                nestedData: false
            },
            gender: {
                isChart: true,
                title: "Gender",
                labels: ['Male', 'Female'],
                dataLabel: 'gender',
                nestedData: false
            },
            Interests: {
                isChart: true,
                title: "Interests",
                dataLabel: 'interestedIn',
                labels: [],
                nestedData: true
            },
            Major: {
                isChart: true,
                title: "Major",
                labels: [],
                dataLabel: 'major',
                nestedData: false
            },
            technology: {
                isChart: true,
                title: "Technology",
                dataLabel: 'knownTech',
                labels: [],
                nestedData: true
            }
        };
        $scope.STARTFILTERS = {
            Status: {
                isToggle: false,
                label: "Status",
                dataLabel: 'status',
                simpleFilter: true,
                nestedHeaders: true
            },
            Grade: {
                isToggle: false,
                label: "Grade",
                dataLabel: 'gradeLabel',
                simpleFilter: true,
                nestedHeaders: true
            },
            GradYear: {
                isToggle: false,
                label: "Graduation Year",
                dataLabel: 'gradYear',
                simpleFilter: true,
                nestedHeaders: true
            },
            University: {
                isToggle: false,
                label: "University",
                dataLabel: 'university',
                simpleFilter: true,
                nestedHeaders: true
            },
            Technologies: {
                isToggle: true,
                label: "Technologies",
                dataLabel: 'knownTech',
                simpleFilter: false,
                nestedHeaders: false
            },
            Interests: {
                isToggle: true,
                label: "Interests",
                dataLabel: 'interestedIn',
                simpleFilter: false,
                nestedHeaders: false
            },
            Major: {
                isToggle: false,
                label: "Major",
                dataLabel: 'major',
                simpleFilter: true,
                nestedHeaders: true
            },
            WorkStatus: {
                isToggle: false,
                label: "Work Status",
                dataLabel: 'workStatus',
                simpleFilter: true,
                nestedHeaders: true
            },
            Name: {
                isToggle: false,
                label: "Name",
                dataLabel: 'name',
                nestedHeaders: true
            }
        };
        $scope.TABLEHEADERS = [
            {
                title: 'Name',
                displayPropertyName: 'name',
                sortPropertyName: 'name',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Major',
                displayPropertyName: 'major',
                sortPropertyName: 'major',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'School',
                displayPropertyName: 'university',
                sortPropertyName: 'university',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Graduation Year',
                displayPropertyName: 'gradYear',
                sortPropertyName: 'gradYear',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Status',
                displayPropertyName: 'status',
                sortPropertyName: 'status',
                sort: 'ascending',
                selected: false
            },
            {
                title: 'Grade',
                displayPropertyName: 'gradeLabel',
                sortPropertyName: 'gradeValue',
                sort: 'descending',
                selected: false
            }];
        $scope.DATA = null;
        $scope.PATH = 'techpoint';
        $scope.isCompany = false;
=======
    .controller('TechPointDashboardCtrl', ['$scope', 'TechPointDashboardService', function($scope, TechPointDashboardService){
    //BEGIN CONFIG DATA
    $scope.STARTCHARTSANDSTATS = {
        University: {
            isChart:false,
            title: "Universities",
            icon:'university',
            dataLabel: 'university',
            nestedData: false
        },
        status:{
            isChart:true,
            title:"Stage",
            labels: ["Stage 1 Approved", "Remaining", "Denied"],
            dataLabel:'status',
            nestedData: false
        },
        gender:{
            isChart:true,
            title:"Gender",
            labels:['male','female'],
            dataLabel:'gender',
            nestedData: false
        },
        //Interests: {
        //    isChart:true,
        //    title: "Interests",
        //    dataLabel: 'interestedIn',
        //    labels: [],
        //    nestedData: true
        //},
        Major: {
            isChart:true,
            title: "Major",
            labels: [],
            dataLabel: 'major',
            nestedData: false
        },
        technology:{
            isChart:true,
            title:"Technology",
            dataLabel:'knownTech',
            labels:[],
            nestedData: true
        }
    };
    $scope.STARTFILTERS = {
        Status: {
            isToggle: false,
            label: "Status",
            dataLabel: 'status',
            simpleFilter:true,
            nestedHeaders:true
        },
        Grade: {
            isToggle: false,
            label: "Grade",
            dataLabel: 'gradeLabel',
            simpleFilter:true,
            nestedHeaders:true
        },
        GradYear: {
            isToggle: false,
            label: "Graduation Year",
            dataLabel: 'gradYear',
            simpleFilter:true,
            nestedHeaders:true
        },
        University: {
            isToggle: false,
            label: "University",
            dataLabel: 'university',
            simpleFilter:true,
            nestedHeaders:true
        },
        Technologies: {
            isToggle: true,
            label: "Technologies",
            dataLabel: 'knownTech',
            simpleFilter:false,
            nestedHeaders:false
        },
        //Interests: {
        //    isToggle: true,
        //    label: "Interests",
        //    dataLabel: 'interestedIn',
        //    simpleFilter:false,
        //    nestedHeaders:true
        //},
        Major: {
            isToggle: false,
            label: "Major",
            dataLabel: 'major',
            simpleFilter:true,
            nestedHeaders:true
        },
        WorkStatus: {
            isToggle: false,
            label: "Work Status",
            dataLabel: 'workStatus',
            simpleFilter:true,
            nestedHeaders:true
        },
        Name: {
            isToggle: false,
            label: "Name",
            dataLabel: 'name',
            simpleFilter:true,
            nestedHeaders:false
        }
    };
    $scope.TABLEHEADERS = [
        {
            title: 'Name',
            displayPropertyName: 'name',
            sortPropertyName: 'name',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'Major',
            displayPropertyName: 'major',
            sortPropertyName: 'major',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'School',
            displayPropertyName: 'university',
            sortPropertyName: 'university',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'Graduation Year',
            displayPropertyName: 'gradYear',
            sortPropertyName: 'gradYear',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'Status',
            displayPropertyName: 'status',
            sortPropertyName: 'status',
            sort: 'ascending',
            selected: false
        },
        {
            title: 'Grade',
            displayPropertyName: 'gradeLabel',
            sortPropertyName: 'gradeValue',
            sort: 'descending',
            selected: false
        }];
    $scope.DATA = null;
    $scope.PATH ='techpoint';
    $scope.isCompany = false;
>>>>>>> master

    }]);
