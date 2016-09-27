/* @flow */
const gulp = require('gulp');
const browserify = require('browserify');
const source = require('vinyl-source-stream');
const buffer = require('vinyl-buffer');
const uglify = require('gulp-uglify');

gulp.task('frontend', function () {

    return browserify({entries: './frontend/init.jsx', debug: true})
        .transform("babelify")
        .bundle()
        .pipe(source('index.js'))
        .pipe(buffer())
        .pipe(uglify())
        .pipe(gulp.dest('./public/js/'))
});

gulp.task('default',['frontend']);
