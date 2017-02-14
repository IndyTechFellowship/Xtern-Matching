module.exports = function(grunt) {

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        uglify: {
            options: {
                mangle: false
            },
            my_target: {
                files: {
                    'output.min.js': ['bundle.js']
                }
            }
        },
        cssmin: {
            options: {
                mergeIntoShorthands: false,
                roundingPrecision: -1
            },
            target: {
                files: {
                    'output.min.css': ['node_modules/angular-centered/angular-centered.css',
                        'node_modules/angular-chart.js/dist/angular-chart.min.css',
                        'node_modules/ng-sortable/dist/ng-sortable.min.css',
                        'node_modules/toastr/build/toastr.min.css',
                        'node_modules/angularjs-slider/dist/rzslider.css',
                        'public/css/studentProfile.css',
                        'public/css/universal.css']
                }
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-cssmin');
    grunt.registerTask('default', ['uglify','cssmin']);
    grunt.registerTask('css', ['cssmin']);

};
