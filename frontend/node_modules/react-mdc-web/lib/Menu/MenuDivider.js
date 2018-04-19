'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _ListDivider = require('../List/ListDivider');

var _ListDivider2 = _interopRequireDefault(_ListDivider);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var MenuDivider = function MenuDivider(_ref) {
  var props = _objectWithoutProperties(_ref, []);

  return _react2.default.createElement(_ListDivider2.default, _extends({
    role: 'menuitem'
  }, props));
};
exports.default = MenuDivider;