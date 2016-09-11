'use strict';

var proxySnippet = require('grunt-connect-proxy/lib/utils').proxyRequest;
var serveStatic = require('serve-static');

module.exports = function (grunt) {
  require('load-grunt-tasks')(grunt);
  require('time-grunt')(grunt);

  grunt.initConfig({
    yeoman: {
      // configurable paths
      app: require('./bower.json').appPath || 'public',
      dist: 'public'
    },
    sync: {
      dist: {
        files: [{
          cwd: '<%= yeoman.app %>',
          dest: '<%= yeoman.dist %>',
          src: '**'
        }]
      }
    },
    watch: {
      options: {
        livereload: 35729
      },
      src: {
        files: [
          '<%= yeoman.app %>/*.html',
          '<%= yeoman.app %>/css/**/*',
          '<%= yeoman.app %>/js/**/*',
          '<%= yeoman.app %>/views/**/*'
        ]
      }
    },
    connect: {
      proxies: [
        {
          context: '/api/v1',
          host: 'localhost',
          port: 5050,
          https: false,
          changeOrigin: false
        }
      ],
      options: {
        port: 5051,
        // Change this to '0.0.0.0' to access the server from outside.
        hostname: 'localhost',
        livereload: 35729
      },
      livereload: {
        options: {
          open: true,
          base: [
            '<%= yeoman.app %>'
          ],
          middleware: function (connect) {
            return [
              proxySnippet,
              serveStatic(require('path').resolve('public'))
            ];
          }
        }
      }
    },
    // Put files not handled in other tasks here
    copy: {
      dist: {
        files: [{
          expand: true,
          dot: true,
          cwd: '<%= yeoman.app %>',
          dest: '<%= yeoman.dist %>',
          src: '**'
        }]
      }
    },
    // Test settings
    karma: {
      unit: {
        configFile: 'test/config/karma.conf.js',
        singleRun: true
      }
    },
    bowercopy: {
      options: {
        destPrefix: '<%= yeoman.app %>'
      },
      test: {
        files: {
          'test/lib/angular-mocks': 'angular-mocks',
          'test/lib/angular-scenario': 'angular-scenario'
        }
      }
    }
  });

  grunt.registerTask('server', function () { //target
    grunt.task.run([
      'configureProxies',
      'connect:livereload',
      'watch'
    ]);
  });
};
