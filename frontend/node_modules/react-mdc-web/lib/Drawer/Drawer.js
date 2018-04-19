'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _Permanent = require('./Permanent');

var _Permanent2 = _interopRequireDefault(_Permanent);

var _Temporary = require('./Temporary');

var _Temporary2 = _interopRequireDefault(_Temporary);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  permanent: _propTypes2.default.bool
};

var Drawer = function Drawer(_ref) {
  var className = _ref.className,
      children = _ref.children,
      permanent = _ref.permanent,
      otherProps = _objectWithoutProperties(_ref, ['className', 'children', 'permanent']);

  return permanent ? _react2.default.createElement(
    _Permanent2.default,
    _extends({
      className: className
    }, otherProps),
    children
  ) : _react2.default.createElement(
    _Temporary2.default,
    _extends({
      className: className
    }, otherProps),
    children
  );
};
Drawer.propTypes = propTypes;
exports.default = Drawer;