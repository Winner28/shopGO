'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});
function isSupportCssCustomProperties() {
  var globalObj = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : window;

  if ('CSS' in globalObj) {
    return globalObj.CSS.supports('(--color: red)');
  }
  return false;
}
exports.default = isSupportCssCustomProperties;