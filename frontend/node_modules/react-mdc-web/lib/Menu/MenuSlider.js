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

var _ScaledComponent = require('./ScaledComponent');

var _ScaledComponent2 = _interopRequireDefault(_ScaledComponent);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var ROOT = 'mdc-simple-menu';
var ANIMATING = ROOT + '--animating';
var OPEN = ROOT + '--open';
var TOPLEFT = OPEN + '-from-top-left';
var TOPRIGHT = OPEN + '-from-top-right';
var BOTTOMLEFT = OPEN + '-from-bottom-left';
var BOTTOMRIGHT = OPEN + '-from-bottom-right';

var propTypes = {
  animation: _propTypes2.default.bool,
  bottom: _propTypes2.default.bool,
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  open: _propTypes2.default.bool,
  right: _propTypes2.default.bool,
  style: _propTypes2.default.object
};

var MenuSlider = function MenuSlider(_ref) {
  var _classnames, _extends2;

  var animation = _ref.animation,
      bottom = _ref.bottom,
      children = _ref.children,
      className = _ref.className,
      open = _ref.open,
      right = _ref.right,
      style = _ref.style,
      otherProps = _objectWithoutProperties(_ref, ['animation', 'bottom', 'children', 'className', 'open', 'right', 'style']);

  var classes = (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, ANIMATING, animation), _defineProperty(_classnames, TOPLEFT, !bottom && !right), _defineProperty(_classnames, TOPRIGHT, !bottom && right), _defineProperty(_classnames, BOTTOMLEFT, bottom && !right), _defineProperty(_classnames, BOTTOMRIGHT, bottom && right), _defineProperty(_classnames, OPEN, open), _classnames), className);
  var horizontal = right ? 'right' : 'left';
  var vertical = bottom ? 'bottom' : 'top';
  return _react2.default.createElement(
    'div',
    _extends({
      className: classes,
      role: 'menuitem',
      style: _extends((_extends2 = {}, _defineProperty(_extends2, horizontal, 0), _defineProperty(_extends2, vertical, 0), _extends2), style)
    }, otherProps),
    children
  );
};
MenuSlider.propTypes = propTypes;
exports.default = (0, _ScaledComponent2.default)(MenuSlider);