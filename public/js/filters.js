//angular.module('Xtern')
//.filter('ngRepeatFinish', function($timeout){
//    //from stack overflow http://stackoverflow.com/questions/15207788/calling-a-function-when-ng-repeat-has-finished
//    return function(data){
//        self = this;
//        var flagProperty = '__finishedRendering__';
//        if(!data[flagProperty]){
//            Object.defineProperty(
//                data,
//                flagProperty,
//                {enumerable:false, configurable:true, writable: false, value:{}});
//            $timeout(function(){
//                delete data[flagProperty];
//                $emit('ngRepeatFinished');
//            },0,false);
//        }
//        return data;
//    };
//});