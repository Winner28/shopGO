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

var _constants = require('./constants');

var _constants2 = _interopRequireDefault(_constants);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  component: _propTypes2.default.string
};

var defaultProps = {
  component: 'div'
};

var Typography = function Typography(_ref) {
  var className = _ref.className,
      component = _ref.component,
      otherProps = _objectWithoutProperties(_ref, ['className', 'component']);

  var classes = (0, _classnames2.default)(_constants2.default, className);
  return _react2.default.createElement(component, _extends({
    className: classes
  }, otherProps));
};
Typography.propTypes = propTypes;
Typography.defaultProps = defaultProps;
exports.default = Typography;