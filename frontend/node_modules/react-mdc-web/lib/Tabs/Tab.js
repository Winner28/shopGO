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
  active: _propTypes2.default.bool,
  className: _propTypes2.default.string,
  component: _propTypes2.default.string
};

var defaultProps = {
  component: 'a'
};

var ROOT = 'mdc-tab';
var ACTIVE = ROOT + '--active';

var Tab = function Tab(_ref) {
  var active = _ref.active,
      className = _ref.className,
      component = _ref.component,
      otherProps = _objectWithoutProperties(_ref, ['active', 'className', 'component']);

  var classes = (0, _classnames3.default)(ROOT, _defineProperty({}, ACTIVE, active), className);
  return _react2.default.createElement(component, _extends({
    className: classes
  }, otherProps));
};
Tab.propTypes = propTypes;
Tab.defaultProps = defaultProps;
exports.default = Tab;