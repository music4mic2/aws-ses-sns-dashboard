var gulp = require('gulp'),
  del = require('del'),
  cssmin = require('gulp-minify-css'),
  browserify = require('browserify'),
  globalShim = require('browserify-global-shim').configure({
    "react": "React",
    "react-dom": "ReactDOM",
  }),
  uglify = require('gulp-uglify'),
  concat = require('gulp-concat'),
  source = require('vinyl-source-stream'),
  buffer = require('vinyl-buffer');

/**
 * Cleaning dist/ folder
 */
gulp.task('clean', function(cb) {
  console.log('clean');
  del(['dist/**'], cb);
})

/**
 * Css compilation
 */
gulp.task('css:min', function() {
  return gulp.src('./app/css/*.css')
  .pipe(concat('style.min.css'))
  .pipe(cssmin())
  .pipe(gulp.dest('./assets/css'));
})

/** JavaScript compilation */
//Js path main
gulp.task('js:min', function() {
  return browserify({entries: './app/js/notification.js', standalone: 'Notification',  extensions: [ '.jsx', '.js' ]})
  .external(["react", "react-dom"])
  .transform('babelify', {presets: ['es2015', 'react']})
  .transform(globalShim)
  .bundle()
  .pipe(source('notification.js'))
  .pipe(buffer())
  .pipe(uglify())
  .pipe(gulp.dest('./assets/js'));
})

/**
 * Compiling resources
 */
gulp.task('build', ['clean', 'css:min', 'js:min'])

gulp.task('default', function() {
  console.log("gulp clean       -> Clean the dist directory");
  console.log("gulp css:min    -> Build the minified styles");
  console.log("gulp js:min      -> Build the minified javascript");
});
