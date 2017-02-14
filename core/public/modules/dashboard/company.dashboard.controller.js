angular.module('Xtern')
    .controller('CompanyDashboardCtrl', ['$scope', 'TechPointDashboardService', function ($scope, TechPointDashboardService) {
        //BEGIN CONFIG DATA
        $scope.STARTCHARTSANDSTATS = {
            University: {
                isChart: false,
                title: "Universities",
                icon: 'university',
                dataLabel: 'university',
                nestedData: false
            },
            gender: {
                isChart: true,
                title: "Gender",
                labels: ['male', 'female'],
                dataLabel: 'gender',
                nestedData: false
            },
            //Interests: {
            //    isChart: true,
            //    title: "Interests",
            //    dataLabel: 'interestedIn',
            //    labels: [],
            //    nestedData: true
            //},
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
            Grade: {
                isToggle: false,
                label: "Grade",
                dataLabel: 'gradeLabel',
                simpleFilter: true,
                nestedHeaders: true,
            },
            GradYear: {
                isToggle: false,
                label: "Graduation Year",
                dataLabel: 'gradYear',
                simpleFilter: true,
                nestedHeaders: true,
            },
            University: {
                isToggle: false,
                label: "University",
                dataLabel: 'university',
                simpleFilter: true,
                nestedHeaders: true,
            },
            Technologies: {
                isToggle: true,
                label: "Technologies",
                dataLabel: 'knownTech',
                simpleFilter: false,
                nestedHeaders: false
            },
            //Interests: {
            //    isToggle: true,
            //    label: "Interests",
            //    dataLabel: 'interestedIn',
            //    simpleFilter: false,
            //    nestedHeaders: false,
            //},
            Major: {
                isToggle: false,
                label: "Major",
                dataLabel: 'major',
                simpleFilter: true,
                nestedHeaders: true,
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
            }];
        $scope.DATA = null;
        $scope.PATH = 'company';
        $scope.isCompany = true;

    }]);
