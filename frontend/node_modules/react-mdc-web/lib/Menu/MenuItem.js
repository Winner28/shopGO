'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _ListItem = require('../List/ListItem');

var _ListItem2 = _interopRequireDefault(_ListItem);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  children: _propTypes2.default.node,
  delay: _propTypes2.default.number
};

var MenuItem = function MenuItem(_ref) {
  var children = _ref.children,
      delay = _ref.delay,
      otherProps = _objectWithoutProperties(_ref, ['children', 'delay']);

  var style = delay ? { transitionDelay: delay.toFixed(3) + 's' } : {};
  return _react2.default.createElement(
    _ListItem2.default,
    _extends({
      role: 'menuitem',
      style: style
    }, otherProps),
    children
  );
};
MenuItem.propTypes = propTypes;
exports.default = MenuItem;