'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _List = require('../List/List');

var _List2 = _interopRequireDefault(_List);

var _ScaledComponent = require('./ScaledComponent');

var _ScaledComponent2 = _interopRequireDefault(_ScaledComponent);

var _DelayedItems = require('./DelayedItems');

var _DelayedItems2 = _interopRequireDefault(_DelayedItems);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  children: _propTypes2.default.node
};

var MenuList = function MenuList(_ref) {
  var children = _ref.children,
      otherProps = _objectWithoutProperties(_ref, ['children']);

  return _react2.default.createElement(
    _List2.default,
    _extends({
      className: 'mdc-simple-menu__items',
      role: 'menu'
    }, otherProps),
    children
  );
};
MenuList.propTypes = propTypes;
exports.default = (0, _ScaledComponent2.default)((0, _DelayedItems2.default)(MenuList));