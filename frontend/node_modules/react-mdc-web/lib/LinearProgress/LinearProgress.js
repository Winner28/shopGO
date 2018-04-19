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

var _transformStyleProperties = require('../utils/transformStyleProperties');

var _transformStyleProperties2 = _interopRequireDefault(_transformStyleProperties);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  accent: _propTypes2.default.bool,
  buffer: _propTypes2.default.number,
  className: _propTypes2.default.string,
  indeterminate: _propTypes2.default.bool,
  progress: _propTypes2.default.number,
  reversed: _propTypes2.default.bool
};

var ROOT = 'mdc-linear-progress';
var INDETERMINATE = ROOT + '--indeterminate';
var ACCENT = ROOT + '--accent';
var REVERSED = ROOT + '--reversed';

var LinearProgress = function LinearProgress(_ref) {
  var _classnames;

  var accent = _ref.accent,
      buffer = _ref.buffer,
      className = _ref.className,
      indeterminate = _ref.indeterminate,
      progress = _ref.progress,
      reversed = _ref.reversed,
      otherProps = _objectWithoutProperties(_ref, ['accent', 'buffer', 'className', 'indeterminate', 'progress', 'reversed']);

  var classes = (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, ACCENT, accent), _defineProperty(_classnames, INDETERMINATE, indeterminate), _defineProperty(_classnames, REVERSED, reversed), _classnames), className);
  var progressScales = {};
  _transformStyleProperties2.default.forEach(function (property) {
    progressScales[property] = 'scaleX(' + progress + ')';
  });

  var bufferScales = {};
  _transformStyleProperties2.default.forEach(function (property) {
    bufferScales[property] = 'scaleX(' + buffer + ')';
  });

  return _react2.default.createElement(
    'div',
    _extends({
      role: 'progressbar',
      className: classes
    }, otherProps),
    _react2.default.createElement('div', { className: 'mdc-linear-progress__buffering-dots' }),
    _react2.default.createElement('div', {
      className: 'mdc-linear-progress__buffer',
      style: _extends({}, bufferScales)
    }),
    _react2.default.createElement(
      'div',
      {
        className: 'mdc-linear-progress__bar mdc-linear-progress__primary-bar',
        style: _extends({}, progressScales)
      },
      _react2.default.createElement('span', { className: 'mdc-linear-progress__bar-inner' })
    ),
    _react2.default.createElement(
      'div',
      { className: 'mdc-linear-progress__bar mdc-linear-progress__secondary-bar' },
      _react2.default.createElement('span', { className: 'mdc-linear-progress__bar-inner' })
    )
  );
};
LinearProgress.propTypes = propTypes;
exports.default = LinearProgress;