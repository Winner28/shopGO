'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _constants = require('./constants');

var _constants2 = _interopRequireDefault(_constants);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var ADJUST_MARGIN = _constants2.default + '--adjust-margin';

var propTypes = {
  component: _propTypes2.default.string,
  adjustMargin: _propTypes2.default.bool
};

var TypographyElement = function TypographyElement(_ref) {
  var adjustMargin = _ref.adjustMargin,
      className = _ref.className,
      component = _ref.component,
      modificator = _ref.modificator,
      otherProps = _objectWithoutProperties(_ref, ['adjustMargin', 'className', 'component', 'modificator']);

  var classes = (0, _classnames3.default)(_constants2.default + '--' + modificator, _defineProperty({}, ADJUST_MARGIN, adjustMargin), className);
  return (0, _react.createElement)(component, _extends({
    className: classes
  }, otherProps));
};

TypographyElement.propTypes = propTypes;

exports.default = TypographyElement;