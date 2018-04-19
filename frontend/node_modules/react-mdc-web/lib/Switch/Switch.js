'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var ROOT = 'mdc-switch';
var NATIVE_CONTROL = ROOT + '__native-control';
var BACKGROUND = ROOT + '__background';
var KNOB = ROOT + '__knob';

var propTypes = {
  className: _propTypes2.default.string,
  disabled: _propTypes2.default.bool
};

var Switch = function Switch(_ref) {
  var className = _ref.className,
      otherProps = _objectWithoutProperties(_ref, ['className']);

  var classes = (0, _classnames2.default)(ROOT, className);

  return _react2.default.createElement(
    'div',
    { className: classes },
    _react2.default.createElement('input', _extends({
      className: NATIVE_CONTROL,
      type: 'checkbox'
    }, otherProps)),
    _react2.default.createElement(
      'div',
      { className: BACKGROUND },
      _react2.default.createElement('div', { className: KNOB })
    )
  );
};
Switch.propTypes = propTypes;
exports.default = Switch;