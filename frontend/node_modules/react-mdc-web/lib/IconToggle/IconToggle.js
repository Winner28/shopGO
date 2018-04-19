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

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  accent: _propTypes2.default.bool,
  ariaLabel: _propTypes2.default.string,
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  disabled: _propTypes2.default.bool,
  primary: _propTypes2.default.bool
};

var ROOT = 'mdc-icon-toggle';

var IconToggle = function IconToggle(_ref) {
  var _classnames;

  var accent = _ref.accent,
      children = _ref.children,
      className = _ref.className,
      disabled = _ref.disabled,
      primary = _ref.primary,
      otherProps = _objectWithoutProperties(_ref, ['accent', 'children', 'className', 'disabled', 'primary']);

  return _react2.default.createElement(
    'i',
    _extends({
      className: (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, ROOT + '--accent', accent), _defineProperty(_classnames, ROOT + '--disabled', disabled), _defineProperty(_classnames, ROOT + '--primary', primary), _classnames), className),
      role: 'button'
    }, otherProps),
    children
  );
};
IconToggle.propTypes = propTypes;
exports.default = IconToggle;