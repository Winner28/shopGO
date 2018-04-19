'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _MenuItem = require('./MenuItem');

var _MenuItem2 = _interopRequireDefault(_MenuItem);

var _MenuDivider = require('./MenuDivider');

var _MenuDivider2 = _interopRequireDefault(_MenuDivider);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var MenuChilds = function MenuChilds(propValue, key, componentName) {
  var type = propValue[key].type;

  if (type !== _MenuItem2.default && type !== _MenuDivider2.default) {
    return new Error('Invalid prop \'children[' + key + ']\' supplied to \'' + componentName + '\', expected \'MenuItem or MenuDivider\'');
  }
  return false;
};
exports.default = MenuChilds;