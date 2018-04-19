'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _isDefined = require('../utils/isDefined');

var _isDefined2 = _interopRequireDefault(_isDefined);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  transition: _propTypes2.default.bool,
  z: _propTypes2.default.number
};

var Elevation = function Elevation(_ref) {
  var _classnames;

  var className = _ref.className,
      children = _ref.children,
      z = _ref.z,
      transition = _ref.transition,
      otherProps = _objectWithoutProperties(_ref, ['className', 'children', 'z', 'transition']);

  return _react2.default.createElement(
    'div',
    _extends({
      className: (0, _classnames3.default)((_classnames = {}, _defineProperty(_classnames, 'mdc-elevation--z' + z, (0, _isDefined2.default)(z)), _defineProperty(_classnames, 'mdc-elevation-transition', transition), _classnames), className)
    }, otherProps),
    children
  );
};
Elevation.propTypes = propTypes;
exports.default = Elevation;