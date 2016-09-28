module.exports = function(grunt) {

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json')
        // uglify: {
        //     options: {
        //         banner: '/*! <%= pkg.name %> <%= grunt.template.today("dd-mm-yyyy") %> */\n'
        //     },
        //     dist: {
        //         files: {
        //             'dist/<%= pkg.name %>.min.js': ['<%= concat.dist.dest %>']
        //         }
        //     }
        // }
        // jshint: {
        //     files: ['Gruntfile.js', 'public/js/*.js'],
        //     options: {
        //         // options here to override JSHint defaults
        //         globals: {
        //             jQuery: true,
        //             console: true,
        //             module: true,
        //             document: true
        //         }
        //     }
    });

    grunt.loadNpmTasks('grunt-contrib-uglify');
    //grunt.loadNpmTasks('grunt-contrib-jshint');
    grunt.registerTask('default', ['jshint', 'uglify']);

};
