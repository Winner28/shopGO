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

var propTypes = {
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  compact: _propTypes2.default.bool,
  dense: _propTypes2.default.bool,
  raised: _propTypes2.default.bool
};

var Button = function Button(_ref) {
  var children = _ref.children,
      className = _ref.className,
      compact = _ref.compact,
      dense = _ref.dense,
      raised = _ref.raised,
      otherProps = _objectWithoutProperties(_ref, ['children', 'className', 'compact', 'dense', 'raised']);

  var classes = (0, _classnames2.default)('mdc-button', {
    'mdc-button--compact': compact,
    'mdc-button--dense': dense,
    'mdc-button--raised': raised
  }, className);
  return _react2.default.createElement(
    'button',
    _extends({
      className: classes
    }, otherProps),
    children
  );
};
Button.propTypes = propTypes;
exports.default = Button;